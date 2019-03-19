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
})


function Load() {
    $.get(domain + "/Projects", function (data, status) {
        $.each(data, function (index, element) {
            $tr = `<tr>
                    <th scope="row">
                        <div class="form-check form-check-inline col-sm-2">
                            <input class="form-check-input" type="checkbox" >
                        </div>
                    </th>
                    <td>`+ element.Name + `</td>
                    <td>`+ element.RouterGroups + `</td>
                    <td>
                        <button id="btn-delete" type="button" class="btn btn-danger">Delete</button>
                        <hidden value=`+ element.ID + `></hidden>
                    </td>
                </tr>`
            $("#tbl-project tbody").append($tr);
        })
    })
}