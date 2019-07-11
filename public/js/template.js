var domain = "http://localhost:7000/api";

$(function () {
    $("#templateForm").validate({
        rules: {
            version: "required",
            proxypass: "required",
            method: {
                required: true,
                minlength: 1
            },
            templateName: "required"
        },
        message: {
            version: "Please enter a version. e.g. 2.0",
            proxypass: "Please enter a proxy pass. e.g. wlxapi:7300",
            method: {
                required: "Please tick method",
                minlength: "You method must be at least 1"
            },
            templateName: "Please enter template name"
        },
    });

    function validateTag(action) {
        if (action == "addTag") {
            $("#tag-error").remove();
        } else {
            if ($("#resgroup").children("span.tag").length == 0) {
                if ($("#tag-error").length == 0) {
                    $tagerr = '<label id="tag-error" for="resgroup" class="error">You must add at least one resource</label>';
                    $("#resgroup").append($tagerr);
                    return false;
                } else {
                    $("#tag-error").remove();
                }
            }
        }
        return true;
    }

    // // Empty input text
    // $("input[text]").focus(function () {
    //     console.log("222");
    //     $(this).val("");
    // })

    // toggle checkbox
    $("#chk-method-any").on("change", function () {
        if ($(this).prop("checked")) {
            $('input[type="checkbox"]').each(function (index) {
                $(this).prop("checked", true);
            })
        } else {
            $('input[type="checkbox"]').each(function (index) {
                $(this).prop("checked", false);
            })
        }
    })

    function addTag(resource) {
        $label = '<span name="restag" class="tag mr-2 badge badge-success" style="font-size:16px">' + resource + '</span>';
        $("#resgroup").append($label);
    }

    // Add resource tag
    $("#btn-resource").click(function () {
        $("#resource").valid({
            rules: {
                resource: "required"
            },
            message: {
                resource: "Please enter a resource name"
            }
        });

        validateTag("addTag");
        resource = $(this).siblings("input").val();

        addTag(resource);
    })



    // Delete resource tag
    $("div").on("click", "[name=restag]", function () {
        $(this).remove();
    })

    // Empty button
    $("#btn-clear").click(function () {
        $("input[text]").val("");
        $("#div-resource").remove();
    })

    // Tick checkbox
    // $(".form-check").click(function () {
    //     $method = $(this).find("input");
    //     if ($method.attr("checked") == "checked") {
    //         $method.removeAttr("checked");
    //     } else {
    //         $method.attr("checked", "checked");
    //     }
    // })

    $("#project-dropdown").click(function () {
        name = $(this).find(":selected").text()
        if (name == "---") {
            return;
        }
        $.get(domain + "/Projects/" + name, function (data, status) {
            $("#router-dropdown").children("option").remove();
            var option = data.RouterGroups.split(",");
            $.each(option, function (index, element) {
                $option = '<option value=' + element + '>' + element + '</option>';
                $("#router-dropdown").append($option);
            })
        })
    })

    // Post template 
    $("#btn-create-template").click(function () {
        if ($("#templateForm").valid() == false) {
            return
        }
        if (validateTag("") == false) {
            return
        }

        var resources = ""
        $(".tag").each(function (index) {
            resources += $(this).html() + ",";
        })
        resources = resources.slice(0, resources.length - 1);

        var methods = ""
        $("[name='method']").each(function (index) {
            console.log($(this).prop("checked"));
            if ($(this).prop("checked") == true) {
                methods += $(this).val() + ",";
            }
        })

        if ($("#chk-method-any").prop("checked") == true) {
            methods = "POST,PATCH,DELETE,GET,PUT,OPTIONS,"
        }

        methods = methods.slice(0, methods.length - 1);
        var template = {
            Resource: resources,
            Method: methods,
            Version: $("#version").val(),
            ProxySchema: $("#proxy-schema-dropdown").find(":selected").text(),
            ProxyPass: $("#proxy-pass").val(),
            ProxyVersion: $("#proxy-version").val(),
            CustomConfig: $("#custom-config-textarea").html(),
            ProjectName: $("#project-dropdown").find(":selected").text(),
            RouterGroup: $("#router-dropdown").find(":selected").text(),
            TemplateName: $("#template").val()
        }

        console.log(template);

        // Post
        $.post(domain + "/Templates", template, function (result) {
            // locaton.reload(true);
            window.location.href = "http://localhost:7000/home";
        })
    })

    // Edit the template
    $("#btn-edit-template").click(function () {
        if ($("#templateForm").valid() == false) {
            return
        }
        if (validateTag("") == false) {
            return
        }

        var resources = ""
        $(".tag").each(function (index) {
            resources += $(this).html() + ",";
        })
        resources = resources.slice(0, resources.length - 1);

        var methods = ""
        $("[name='method']").each(function (index) {
            if ($(this).prop("checked") == true) {
                methods += $(this).val() + ",";
            }
        })

        if ($("#chk-method-any").prop("checked") == true) {
            methods = "POST,PATCH,DELETE,GET,PUT,OPTIONS,"
        }

        methods = methods.slice(0, methods.length - 1);
        url = $(location).attr("href");
        params = url.split("?")[1];
        id = params.split("=")[1];
        var template = {
            ID: id,
            Resource: resources,
            Method: methods,
            Version: $("#version").val(),
            ProxySchema: $("#proxy-schema-dropdown").find(":selected").text(),
            ProxyPass: $("#proxy-pass").val(),
            ProxyVersion: $("#proxy-version").val(),
            CustomConfig: $("#custom-config-textarea").html(),
            ProjectName: $("#project-dropdown").find(":selected").text(),
            RouterGroup: $("#router-dropdown").find(":selected").text(),
            TemplateName: $("#template").val()
        }

        console.log(template);
        // Patch
        $.ajax({
            url: domain + "/Templates/" + id,
            type: 'PATCH',
            data: JSON.stringify(template),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (result) {
                $("#resgroup").children().remove();
                // Load();
                window.location.href = "http://localhost:7000/home";
            },
            error: function (request, msg, error) {
                data = request.responseJSON;
                window.location.href = "http://localhost:7000/error?Error=" + data.Error + "&StatusCode=" + data.StatusCode;
            }
        });
    })
    // Delete Project
    $("tbody").on("click", ".btn-delete", function () {
        id = $(this).next().attr("value");
        $.ajax({
            url: domain + "/Projects/" + id,
            method: 'DELETE',
            success: function (result) {
                Load();
            },
            error: function (request, msg, error) {
                data = request.responseJSON;
                window.location.href = "http://localhost:7000/error?Error=" + data.Error + "&StatusCode=" + data.StatusCode;
            }
        });
    })
})


function Load() {
    $.get(domain + "/Projects", function (data, status) {
        $.each(data, function (index, element) {
            $option = '<option class="project" value=' + element.Name + '>' + element.Name + '</option>';
            $("#project-dropdown").append($option);
        })
    })

    url = $(location).attr("href");
    if (url.indexOf('?') != -1) {
        params = url.split("?")[1];
        id = params.split("=")[1];
        if (id != "") {
            $.get(domain + "/Templates/" + id, function (data, status) {
                resources = data.Resource.split(",");
                $.each(resources, function (index, element) {
                    $label = '<span name="resgroup" class="tag mr-2 mt-2 badge badge-success" style="font-size:14px;">' + element + '</span>';
                    $("#resgroup").append($label);
                });

                if (data.Method != "") {
                    methods = data.Method.split(",")
                    $.each(methods, function (index, element) {
                        if (element == "POST") {
                            $("#chk-method-post").attr("checked", "checked");
                        }
                        if (element == "PATCH") {
                            $("#chk-method-patch").attr("checked", "checked");
                        }
                        if (element == "DELETE") {
                            $("#chk-method-delete").attr("checked", "checked");
                        }
                        if (element == "GET") {
                            $("#chk-method-get").attr("checked", "checked");
                        }
                        if (element == "PUT") {
                            $("#chk-method-put").attr("checked", "checked");
                        }
                        if (element == "OPTIONS") {
                            $("#chk-method-option").attr("checked", "checked");
                        }
                        if (element == "ANY") {
                            $("#chk-method-post").attr("checked", "checked");
                            $("#chk-method-patch").attr("checked", "checked");
                            $("#chk-method-delete").attr("checked", "checked");
                            $("#chk-method-get").attr("checked", "checked");
                            $("#chk-method-put").attr("checked", "checked");
                            $("#chk-method-option").attr("checked", "checked");
                        }
                    });
                }

                $("#version").val(data.Version);
                $("#proxy-schema-dropdown option[value=" + data.ProxySchema + "]").prop("selected", true);
                $("#proxy-pass").val(data.ProxyPass);
                $("#proxy-version").val(data.ProxyVersion);
                $("#custom-config-textarea").val(data.CustomConfig);
                $("#project-dropdown option[value=" + data.ProjectName + "]").prop("selected", true);
                $("#template").val(data.TemplateName);
            }).done(function (data) {
                if ($(".project").length) {
                    $.get(domain + "/Projects/" + data.ProjectName, function (data, status) {
                        $("#router-dropdown").children("option").remove();
                        var option = data.RouterGroups.split(",");
                        $.each(option, function (index, element) {
                            $option = '<option value=' + element + '>' + element + '</option>';
                            $("#router-dropdown").append($option);
                        })
                    }).done(function () {
                        $("#router-dropdown option[value=" + data.RouterGroup + "]").prop("selected", true);
                    })
                }
            })
        }
    }
}





