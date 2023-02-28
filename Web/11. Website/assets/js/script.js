document.querySelector(".icon-menu").addEventListener("click", function() { 
    document.querySelector(".nav-left").classList.toggle("visible")
})

document.querySelectorAll(".clickable").forEach(element => {
    element.addEventListener("click", function() { 
            console.log("bonjour")
            document.querySelector(".nav-left").classList.remove("visible")
    })
});

// On switch from phone view to big screen view, removes the "visible" class on 
// the menu (needed only when using burger menu)
window.addEventListener("resize", function() {
    if (this.innerWidth > 950) {
        this.document.querySelector(".nav-left").classList.remove("visible")
    } 
})


if (document.getElementById("form")) {
    document.getElementById("form").addEventListener("submit", function(event) {
        event.preventDefault()
        alert("Noppers")
    })
}

document.querySelector(".hz-menu-icon").addEventListener("click", function() { 
    document.querySelector(".hz-nav").classList.toggle("visible")
})

// document.querySelectorAll(".nav-link").forEach(element => {
//     element.addEventListener("click", function() { 
//         if (window.innerWidth < 750) {
//             document.querySelector(".nav-left").classList.remove("visible")
//         }
//     })
// });

function scrollToDiv(element, offset=20) {
    element.scrollIntoView({block: "center"},true);

}
var scrollAllowed = true
if (document.querySelector(".timeline-container")) {
    document.querySelector("#hz-btn-top").addEventListener("click", function() {
        scrollToDiv(document.querySelector("#tl-2022"))
        removeHzVisible()
    })
    document.querySelector("#hz-btn-2022").addEventListener("click", function() {
        scrollToDiv(document.querySelector("#tl-2022"))
        removeHzVisible()
    })
    document.querySelector("#hz-btn-2023").addEventListener("click", function() {
        scrollToDiv(document.querySelector("#tl-2023"))
        removeHzVisible()
    })
    document.querySelector("#hz-btn-bottom").addEventListener("click", function() {
        scrollToDiv(document.querySelector("#tl-end"))
        removeHzVisible()
    })
}

function removeHzVisible() {
        document.querySelector(".hz-nav").classList.remove("visible")
}