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
    $("#btn-create-project").click(function () {
        var tags = "";
        $(".router-label").each(function (index) {
            tags += $(this).html() + ",";
        })
        var project = {
            Name: $("#input-project").val(),
            RouterGroups: tags.slice(0, tags.length - 1)
        }
        $.post(domain + "/Projects", project, function (result) {
            Load();
        })
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
}