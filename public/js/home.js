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
})


function Load() {
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
                    <td>`+ element.Name+ `</td>
                    <td>
                        <button type="button" class="btn btn-danger btn-delete">Delete</button>
                        <button type="button" class="btn btn-info btn-preview">Preview</button>
                        <button type="button" class="btn btn-dark btn-download">Download</button>
                        <hidden value=`+ element.URL + `></hidden>
                    </td>
                </tr>`
            $("#tbl-file tbody").append($tr);
        })
    })
}