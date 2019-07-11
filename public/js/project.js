var domain = "http://localhost:7000/api";
$(function () {
    $("#projectForm").validate({
        rules: {
            project: "required"
        },
        message: {
            project: "Please enter a project name. e.g. wlxapi"
        },
    });

    function validateTag(action) {
        if (action == "addTag") {
            $("#tag-error").remove();
        } else {
            if ($("#routergroup").children("span.tag").length == 0) {
                if ($("#tag-error").length == 0) {
                    $tagerr = '<label id="tag-error" for="routergroup" class="error">You must add at least one router</label>';
                    $("#routergroup").append($tagerr);
                } else {
                    $("#tag-error").remove();
                }
            }
        }
    }

    // Empty input text
    $(".form-control").focus(function () {
        $(this).val("");
    })

    // Add router tag
    $("#btn-router-group").click(function () {
        $("#group").valid({
            rules: {
                group: "required"
            },
            message: {
                group: "Please enter a group. e.g. api"
            }
        });

        validateTag("addTag");
        router = $(this).siblings("input").val();
        $label = '<span name="routertag" class="tag mr-2 badge badge-success" style="font-size:16px">' + router + '</span>';
        $("#routergroup").append($label);
    })

    // Delete router tag
    $("div").on("click", "[name=routertag]", function () {
        $(this).remove();
    })

    // Empty button
    $("#btn-clear").click(function () {
        $(".form-control").val("");
        $("#div-router-group").remove();
    })

    // Checkbox event
    $("tbody").on("click", ":checkbox", function () {
        if ($(this).is(":checked")) {
            id = $(this).next().attr("value");
            console.log(id);
            $.get(domain + "/Projects/" + id, function (data, status) {
                $("#project").val(data.Name);
                tags = data.RouterGroups.split(",");
                $.each(tags, function (index, element) {
                    $label = '<span name="routertag" class="tag mr-2 badge badge-success" style="font-size:16px">' + element + '</span>';
                    $("#routergroup").append($label);
                })
            });
        } else {
            $("#project").val("");
            $(".tag").remove();
        }
    })


    // Post project 
    $("#btn-create-project").click(function () {
        if ($("#projectForm").valid() == false) {
            return
        }
        validateTag("");

        var tags = "";
        $(".tag").each(function (index) {
            tags += $(this).html() + ",";
        })
        var project = {
            Name: $("#project").val(),
            RouterGroups: tags.slice(0, tags.length - 1)
        }
        $.post(domain + "/Projects", project, function (result) {
            Load();
        })
    })

    // Edit project 
    $("#btn-edit-project").click(function () {
        if ($("#projectForm").valid() == false) {
            return
        }
        validateTag("");

        var tags = "";
        $(".tag").each(function (index) {
            tags += $(this).html() + ",";
        })

        id = $("input:checkbox:checked").next().attr("value");
        var project = {
            ID: id,
            Name: $("#project").val(),
            RouterGroups: tags.slice(0, tags.length - 1)
        }

        // Patch
        $.ajax({
            url: domain + "/Projects/" + id,
            type: 'PATCH',
            data: JSON.stringify(project),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (result) {
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
        $("table > tbody").children().remove();
        $.each(data, function (index, element) {
            $tr = `<tr>
                    <td scope="row">
                        <div class="form-check form-check-inline col-sm-2">
                            <input class="form-check-input" type="checkbox">
                            <hidden value=`+ element.ID + `></hidden>
                        </div>
                    </td>
                    <td>`+ element.Name + `</td>
                    <td>`+ element.RouterGroups + `</td>
                    <td>
                        <button type="button">Delete</button>
                        <hidden value=`+ element.ID + `></hidden>
                    </td>
                </tr>`
            $("#tbl-project tbody").append($tr);
        })
    })
}