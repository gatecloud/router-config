$(function () {
})


function Load() {
    $.get(domain + "/Templates", function (data, status) {
        $.each(data, function (index, element) {
            $tr = `<th scope="row">
                        <div class="form-check form-check-inline col-sm-2">
                            <input class="form-check-input" type="checkbox" >
                        </div>
                    </th>
                    <td>`+ element.ProjectName + `</td>
                    <td>`+ element.GroupName + `</td>
                    <td>`+ element.TemplateName + `</td>
                </tr>`
            $("#tbl-template tbody").append($tr);
        })
    })
}