
function changeDisplayMode() {
    if ($("body").hasClass("theme-dark")==true){
        $("body").removeClass("theme-dark");
    } else {
        $("body").addClass("theme-dark");

    }

    $.get( "/dashboard/swap-dark-mode/", function( data ) {});


    $.ajax({
        type: "POST",
        url: "/dashboard/swap-dark-mode/"
    });
}

