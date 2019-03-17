$(function () {


    $("input[text]").focus(function () {
        $(this).val("");
    })

    $("#btn-resource").click(function () {
        resource = $(this).siblings("input").val();
        $label = '<span class="resource-label badge badge-success">' + resource + '</span>';
        $(this).next("div").append($label);
        alert($(this));

    })


    $("td").on("click", ".resource-label", function () {
        $(this).remove();
    })

    $(".btn-create-project").click(function () {
        $.post(domain + "/Projects", {
            Name: $("#input-project").val()

        }, function (result) {
            alert(result);
        })
    })

    $("#btn-new-template").click(function(){
        if ($(this).prev("input[text]").val() == ""){
            alert("Input template name");
        }
    })

})

var domain = "http://localhost:7000/api";

function LoadProjects() {
    $.get(domain + "/Projects", function (data, status) {
        $.each(data, function (index, element) {
            $option = '<button class="dropdown-item" type="button">' + element.Name + '</button>';
            $("#project-dropdown").append($option);
        })
    })
}