package logic

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileGroup struct {
	URL      string
	FileName string
}

func ToList() ([]FileGroup, error) {
	var groups []FileGroup
	if err := filepath.Walk("groups/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			_, filename := filepath.Split(path)
			group := FileGroup{
				URL:      path,
				FileName: filename,
			}
			groups = append(groups, group)
		}

		return nil
	}); err != nil {
		return nil, nil
	}

	return groups, nil
}

func Load(filename string) (RouterTemplate, error) {
	if filename == "" {
		return RouterTemplate{}, errors.New("filename can not be empty")
	}
	var routerTemplate RouterTemplate
	if err := filepath.Walk("groups/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			_, fname := filepath.Split(path)
			if fname == filename {
				body, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}

				err = json.Unmarshal(body, &routerTemplate)
				if err != nil {
					return err
				}

				return nil
			}
		}
		return nil
	}); err != nil {
		return RouterTemplate{}, nil
	}
	return routerTemplate, nil
}

func Export(filename string, fileList []string) error {
	if filename == "" {
		return errors.New("filename can not be empty")
	}
	var routerTemplates []RouterTemplate
	if err := filepath.Walk("groups/", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			_, fname := filepath.Split(path)
			for _, v := range fileList {
				if strings.TrimSpace(v) == fname {
					var rt RouterTemplate
					body, err := ioutil.ReadFile(path)
					if err != nil {
						return err
					}

					err = json.Unmarshal(body, &rt)
					if err != nil {
						return err
					}
					routerTemplates = append(routerTemplates, rt)
					break
				}
			}
		}
		return nil
	}); err != nil {
		return nil
	}

	var servers []Server
	for _, template := range routerTemplates {
		for _, resource := range template.Resources {
			for _, method := range template.Methods {
				path := ""
				if template.Version != "" {
					path = "/" + template.Version + "/" + resource
				} else {
					path = "/" + resource
				}

				if method == "GETBYID" {
					method = "GET"
					path += "/:id"
				} else if method == "PATCHBYID" {
					method = "PATCH"
					path += "/:id"
				} else if method == "DELETEBYID" {
					method = "DELETE"
					path += "/:id"
				}

				server := Server{
					Method:        method,
					Path:          path,
					ProxyScheme:   template.ProxySchema,
					ProxyPass:     template.ProxyPass,
					ProxyPassPath: "/" + resource,
					APIVersion:    template.ProxyVersion,
					CustomConfigs: template.CustomConfigs,
				}

				servers = append(servers, server)
			}
		}
	}

	return createJSONFile(filename, servers)
}

func createJSONFile(filename string, servers []Server) error {
	var (
		file *os.File
		path string
	)

	folderpath, err := os.Getwd()
	if err != nil {
		return err
	}
	path = folderpath + "\\files\\"

	// detect if file exists
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir("files", 0755)
		if err != nil {
			return err
		}
	}

	match := regexp.MustCompile(`[.\d]*.json\z`).MatchString
	if !match(filename) {
		path = path + filename + ".json"
	}

	file, err = os.Create(path)
	defer file.Close()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(servers, "", "\t")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}
