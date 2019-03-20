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
                alert("good")
            }
        })
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
})


function Load() {
    $("tr").remove();
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
                    <td></td>
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
                        <button type="button" class="btn btn-dark btn-download">Download</button>
                    </td>
                </tr>`
            $("#tbl-file tbody").append($tr);
        })
    })
}