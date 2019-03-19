$(function () {
    $("input[text]").focus(function () {
        $(this).val("");
    })

    $("#btn-resource").click(function () {
        resource = $(this).siblings("input").val();
        $label = '<span class="resource-label badge badge-success">' + resource + '</span>';
        $(this).next("div").append($label);
    })


    $("td").on("click", ".resource-label", function () {
        $(this).remove();
    })

    $(".btn-create-project").click(function () {
        var project = {
            Name: $("#input-project").val()
        }
        $.post(domain + "/Projects", project, function (result) {
            alert(result);
        })
    })

    $(".form-check").click(function () {
        $method = $(this).find("input");
        if ($method.attr("checked") == "checked") {
            $method.removeAttr("checked");
        } else {
            $method.attr("checked", "checked");
        }
    })

    $("#btn-new-template").click(function () {
        // Validation
        if ($(".tm-check").val() == "") {
            alert("No blank");
            return;
        }

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
        methods = methods.slice(0, methods.length - 1);
        var template = {
            Resource: resources,
            Method: methods,
            Version: $("#version").val(),
            ProxySchema: $("#proxySchema").val(),
            ProxyPass: $("#proxyPass").val(),
            ProxyVersion: $("#proxyVersion").val(),
            CustomConfig: $("#custom-config-textarea").html(),
            ProjectName: $("#project-dropdown").find(":selected").text(),
            TemplateName: $("#text-template").val()

        }
        // Post
        $.post(domain + "/Templates", template, function (result) {
            locaton.reload(true);
        })
    })


})

var domain = "http://localhost:7000/api";

function Load() {
    $.get(domain + "/Projects", function (data, status) {
        $.each(data, function (index, element) {
            $option = '<option value=' + element.Name + '>' + element.Name + '</option>';
            $("#project-dropdown").append($option);
        })
    })

    $.get(domain + "/Templates", function (data, status) {
        $.each(data, function (index, element) {
            $template = '<div class="template-group list-group-item list-group-item-action">' +
                '<input type="checkbox" aria-label="Checkbox for following text input">' +
                '<a href= ' + element.URL + '> ' + element.ProjectName+'/'+element.TemplateName + '</a>' +
                '</div>'
            $(".template-list").append($template);
        })
    })

    
}