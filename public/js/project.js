var domain = "http://localhost:7000/api";
$(function () {
    // Empty input text
    $(".form-control").focus(function () {
        $(this).val("");
    })

    // Add router tag
    $("#btn-router-group").click(function () {
        router = $("#input-router").val();
        $label = '<span class="router-label mr-2 badge badge-success">' + router + '</span>';
        $("#div-router-group").append($label);
    })


    // Delete router tag
    $("div").on("click", ".router-label", function () {
        $(this).remove();
    })

    // Empty button
    $("#btn-clear").click(function () {
        $(".form-control").val("");
        $("#div-router-group").remove();
    })

    // Checkbox event
    $("td").on("click", "input:checkbox", function () {
        if ($(this).is(":checked")) {
            id = $(this).next().attr("value");
            console.log(id);
            $.get(domain + "/Projects/" + id, function (data, status) {
                $("#input-project").val(data.Name);
                tags = data.RouterGroups.split(",");
                $.each(tags, function (index, element) {
                    $label = '<span class="router-label mr-2 badge badge-success">' + element + '</span>';
                    $("#div-router-group").append($label);
                })
            });

        } else {
            $("#input-project").val("");
            $("tr").remove();
        }
    })


    // Post project 
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

    // Edit project 
    $("#btn-edit-project").click(function () {
        var tags = "";
        $(".router-label").each(function (index) {
            tags += $(this).html() + ",";
        })

        id = $("input:checkbox:checked").next().attr("value");
        var project = {
            ID: id,
            Name: $("#input-project").val(),
            RouterGroups: tags.slice(0, tags.length - 1)
        }
        // Patch
        $.ajax({
            url: domain + "/Projects",
            type: 'PATCH',
            data: JSON.stringify(project),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function (result) {
                Load();
            },
            error: function (request, msg, error) {
                data = request.responseJSON;
                window.location.href="http://localhost:7000/error?Error="+data.Error+"&StatusCode="+data.StatusCode;
            }
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
                data = request.responseJSON;
                window.location.href="http://localhost:7000/error?Error="+data.Error+"&StatusCode="+data.StatusCode;
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
                        <button type="button" class="btn btn-danger btn-delete">Delete</button>
                        <hidden value=`+ element.ID + `></hidden>
                    </td>
                </tr>`
            $("#tbl-project tbody").append($tr);
        })
    })
}