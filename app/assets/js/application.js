 
import "bootstrap/dist/js/bootstrap.bundle.js";
import "@fortawesome/fontawesome-free/js/all.js";

$(() => {

    $(".search-category div span").on("click", function(){
        var dataCategory = $(this).attr("data-category");
        $("#course-Category").val(dataCategory);
        $(".search-category .dropdown-toggle").html("Search by: " + dataCategory +" ");
    });

    $(".fa-search").on("click", function(){
        $(this).hide("fast");
        $(".fa-times").show("fast");
        $(".search-category").show("fast");
        console.log("click search")
    })

    $(".fa-times").on("click", function(){
        $(this).hide("fast");
        $(".fa-search").show("fast");
        $(".search-category").hide("fast");
        console.log("click times")
    })

});
