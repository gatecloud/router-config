var domain = "http://localhost:7000/api";

$(function () {
    // $.validator.addMethod("quantity", function (value, element) {
    //     alert("111");
    //     console.log($(element).children("span").length);
    //     // return !this.optional(element) && !this.optional($(element).parent().prev().children("select")[0]);
    //     return $(element).children("span").length == 0
    // }, "Please select both the item and its amount.");

    $("#templateForm").validate({
        rules: {
            // resgroup:{
            //     quantity: true
            // },
            version: "required",
            proxypass: "required",
            method: {
                required: true,
                minlength: 1
            }
        },
        message: {
            version: "Please enter a version. e.g. 2.0",
            proxypass: "Please enter a proxy pass. e.g. wlxapi:7300",
            method: {
                required: "Please tick method",
                minlength: "You method must be at least 1"
            }
        },
    });

    function validateTag(action) {
        if (action == "addTag") {
            $("#tag-error").remove();
        } else {
            if ($("#resgroup").children("label.tag").length == 0) {
                if ($("#tag-error").length == 0) {
                    $tagerr = '<label id="tag-error" for="resgroup" class="error">You must add at least one resource</label>';
                    $("#resgroup").append($tagerr);
                } else {
                    $("#tag-error").remove();
                }
            }
        }

    }

    // // Empty input text
    // $("input[text]").focus(function () {
    //     console.log("222");
    //     $(this).val("");
    // })


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
        if (resource == "") {
            console.log("No blank");
            // alert("No blank");
            // return;
        }
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
        $("#templateForm").valid();
        validateTag("");

        var resources = ""
        $(".resource-label").each(function (index) {
            resources += $(this).html() + ",";
        })
        resources = resources.slice(0, resources.length - 1);
        var rsValidation = resources.replace(',', '');
        if (rsValidation == "") {
            console.log("no blank");
            // alert("No blank");
            // return;
        }

        var methods = ""
        $("[name='method']").each(function (index) {
            if ($(this).find("input").attr("checked") == "checked") {
                methods += $(this).find("input").val() + ",";
            }
        })

        if ($("#chk-method-any").attr("checked") == "checked") {
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
            TemplateName: $("#text-template").val()
        }


        // Post
        $.post(domain + "/Templates", template, function (result) {
            // locaton.reload(true);
        })
    })

    // Edit the template
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
                $("#div-resource").children().remove();
                Load();
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

            $("#version").val(data.version);
            $("#proxy-schema-dropdown option[value=" + data.proxyschema + "]").prop("selected", true);
            $("#proxy-pass").val(data.proxypass);
            $("#proxy-version").val(data.proxyversion);
            $("#custom-config-textarea").val(data.customconfig);
            $("#project-dropdown option[value=" + data.ProjectName + "]").prop("selected", true);
            $.get(domain + "/Projects/" + name, function (data, status) {
                $("#router-dropdown").children("option").remove();
                var option = data.RouterGroups.split(",");
                $.each(option, function (index, element) {
                    $option = '<option value=' + element + '>' + element + '</option>';
                    $("#router-dropdown").append($option);
                })
            }).done(function () {
                $("#router-dropdown option[value=" + data.RouterGroup + "]").prop("selected", true);
            })
            $("#text-template").val(data.TemplateName);
        })
    }


}





