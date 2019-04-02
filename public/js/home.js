var domain = "http://localhost:7000/api";
$(function () {

    // Tick checkbox
    $("div").on("click", ".form-check-input", function () {
        if ($(this).attr("checked") == "checked") {
            $(this).removeAttr("checked");
        } else {
            $(this).attr("checked", "checked");
        }
    })

    // Merge template
    $("#btn-merge-template").click(function () {
        var templates = [];
        $(".form-check-input:checkbox:checked").each(function (index) {
            id = $(this).next().attr("value")
            templates.push({
                ID: id
            })
        })

        var file = new Object();
        file.Name = $("#input-file").val();
        file.Templates = templates
        $.ajax({
            url: domain + "/Files",
            type: "POST",
            data: JSON.stringify(file),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function () {
                Load();
            }
        })
    })

    // Delete template file
    $("tbody").on("click", ".btn-delete-template", function () {
        id = $(this).next().attr("value");
        $.ajax({
            url: domain + "/Templates/" + id,
            type: 'DELETE',
            success: function (result) {
                Load();
            },
            error: function (request, msg, error) {
                console.log(request);
                console.log(msg);
                console.log(error);
            }
        });
    })

    // Edit template file
    $("tbody").on("click", ".btn-edit", function () {
        id = $(this).prev().attr("value");
        $(location).attr("href", "http://localhost:7000/template?id=" + id);
    })

    // Delete downloaded file 
    $("tbody").on("click", ".btn-delete", function () {
        id = $(this).next().attr("value");
        $.ajax({
            url: domain + "/Files/" + id,
            type: 'DELETE',
            success: function (result) {
                Load();
            },
            error: function (request, msg, error) {
                console.log(request);
                console.log(msg);
                console.log(error);
            }
        });
    })

    // Preview the merged file
    $("tbody").on("click", ".btn-preview", function () {
        id = $(this).prev().attr("value");
        $(".textarea-preview").empty();
        $.get(domain + "/Files/" + id, function (data, status) {
            console.log(data.Preview);
            $(".textarea-preview").append(data.Preview)
        })
    })
})


function Load() {
    $("table > tbody").children().remove();
    $.get(domain + "/Templates", function (data, status) {
        $.each(data, function (index, element) {
            $tr = `<tr>
                    <th scope="row">
                        <div class="form-check form-check-inline col-sm-2">
                            <input class="form-check-input" type="checkbox" >
                            <hidden value="`+ element.ID + `"></hidden>
                        </div>
                    </th>
                    <td>`+ element.ProjectName + `</td>
                    <td>`+ element.RouterGroup + `</td>
                    <td>`+ element.TemplateName + `</td>
                    <td>
                        <button type="button" class="btn btn-danger btn-delete-template">Delete</button>
                        <hidden value=`+ element.ID + `></hidden>
                        <button type="button" class="btn btn-danger btn-edit">Edit</button>
                    </td>
                </tr>`
            $("#tbl-template tbody").append($tr);
        })
    })
    $.get(domain + "/Files", function (data, status) {
        $.each(data, function (index, element) {
            $tr = `<tr>
                    <td>`+ element.Name + `</td>
                    <td>
                        <button type="button" class="btn btn-danger btn-delete">Delete</button>
                        <hidden value=`+ element.ID + `></hidden>
                        <button type="button" class="btn btn-info btn-preview">Preview</button>
                        <hidden value=`+ element.URL + `></hidden>
                        <a href="`+ element.URL + `" class="btn btn-dark btn-download" >Download</a>
                    </td>
                </tr>`
            $("#tbl-file tbody").append($tr);
        })
    })
}