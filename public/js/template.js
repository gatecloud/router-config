var domain = "http://localhost:7000/api";
$(function () {
    // Empty input text
    $("input[text]").focus(function () {
        $(this).val("");
    })

    // Add resource tag
    $("#btn-resource").click(function () {
        resource = $(this).siblings("input").val();
        $label = '<span class="resource-label mr-2 badge badge-success">' + resource + '</span>';
        $("#div-resource").append($label);
    })


    // Delete resource tag
    $("div").on("click", ".resource-label", function () {
        $(this).remove();
    })

    // Empty button
    $("#btn-clear").click(function () {
        $("input[text]").val("");
        $("#div-resource").remove();
    })

    // Tick checkbox
    $(".form-check").click(function () {
        $method = $(this).find("input");
        if ($method.attr("checked") == "checked") {
            $method.removeAttr("checked");
        } else {
            $method.attr("checked", "checked");
        }
    })

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
        // Validation
        // if ($(".tm-check").val() == "") {
        //     alert("No blank");
        //     return;
        // }

        var resources = ""
        $(".resource-label").each(function (index) {
            resources += $(this).html() + ",";
        })
        resources = resources.slice(0, resources.length - 1);
        var rsValidation = resources.replace(',', '');
        if (rsValidation == "") {
            alert("No blank");
            return;
        }

        var methods = ""
        $("#method-check .form-check").each(function (index) {
            if ($(this).find("input").attr("checked") == "checked") {
                methods += $(this).find("input").val() + ",";
            }
        })

        if ($("#chk-method-any").attr("checked") == "checked") {
            method = "POST,PATCH,DELETE,GET,OPTIONS,"
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
            TemplateName: $("#text-template").val()

        }
        // Post
        $.post(domain + "/Templates", template, function (result) {
            // locaton.reload(true);
        })
    })


    $("#btn-edit-template").click(function () {
        var resources = ""
        $(".resource-label").each(function (index) {
            resources += $(this).html() + ",";
        })
        resources = resources.slice(0, resources.length - 1);
        var rsValidation = resources.replace(',', '');
        if (rsValidation == "") {
            alert("No blank");
            return;
        }

        var methods = ""
        $("#method-check .form-check").each(function (index) {
            if ($(this).find("input").attr("checked") == "checked") {
                methods += $(this).find("input").val() + ",";
            }
        })

        if ($("#chk-method-any").attr("checked") == "checked") {
            method = "POST,PATCH,DELETE,GET,OPTIONS,"
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
            TemplateName: $("#text-template").val()

        }

        // Patch
        $.ajax({
            url: domain + "/Templates/" + id,
            type: 'PATCH',
            data: JSON.stringify(template),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (result) {
                Load();
            },
            error: function (request, msg, error) {
                alert(error);
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
                alert(error);
            }
        });
    })
})


function Load() {
    $.get(domain + "/Projects", function (data, status) {
        $.each(data, function (index, element) {
            $option = '<option value=' + element.Name + '>' + element.Name + '</option>';
            $("#project-dropdown").append($option);
        })
    })


    url = $(location).attr("href");
    params = url.split("?")[1];
    id = params.split("=")[1];
    if (id != "") {
        $.get(domain + "/Templates/" + id, function (data, status) {
            resources = data.Resource.split(",");
            $.each(resources, function (index, element) {
                $label = '<span class="resource-label mr-2 badge badge-success">' + element + '</span>';
                $("#div-resource").append($label);

            });

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
                if (element == "OPTIONS") {
                    $("#chk-method-option").attr("checked", "checked");
                }
                if (element == "ANY") {
                    $("#chk-method-post").attr("checked", "checked");
                    $("#chk-method-patch").attr("checked", "checked");
                    $("#chk-method-delete").attr("checked", "checked");
                    $("#chk-method-get").attr("checked", "checked");
                    $("#chk-method-option").attr("checked", "checked");
                }
            });

            $("#version").val(data.Version);
            // $("#proxy-schema-dropdown option[value=" + data.ProxySchema + "]").prop("selected", true);
            $("#proxy-schema-dropdown option[value=https]").prop("selected", true);
            $("#proxy-pass").val(data.ProxyPass);
            $("#proxy-version").val(data.ProxyVersion);
            $("#custom-config-textarea").val(data.CustomConfig);
            $("#project-dropdown option[value=" + data.ProjectName + "]").prop("selected", true);
            $.get(domain + "/Projects/" + name, function (data, status) {
                $("#router-dropdown").children("option").remove();
                var option = data.RouterGroups.split(",");
                $.each(option, function (index, element) {
                    $option = '<option value=' + element + '>' + element + '</option>';
                    $("#router-dropdown").append($option);
                })
            })
            $("#router-dropdown option[value=" + data.RouterGroup + "]").prop("selected", true);
            $("#text-template").val(data.TemplateName);
        })
    }


}