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

})