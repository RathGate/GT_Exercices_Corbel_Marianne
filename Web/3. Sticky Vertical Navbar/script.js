document.querySelector(".icon-menu").addEventListener("click", function() { 
    document.querySelector("header").classList.toggle("visible")
    arrowToggle(document.querySelector(".arrow-icon"))
})

document.querySelectorAll(".nav-link").forEach(element => {
    element.addEventListener("click", function() { 
        if (window.innerWidth < 750) {
            document.querySelector("header").classList.remove("visible")
        }
    })
});

// On switch from phone view to big screen view, removes the "visible" class on 
// the menu (needed only when using burger menu)
window.addEventListener("resize", function() {
    if (this.innerWidth > 950) {
        this.document.querySelector("header").classList.remove("visible")
        arrowClose(this.document.querySelector(".arrow-icon"))
    } 
})

function arrowToggle(element) {
    if (element) {
        if (element.innerHTML == "navigate_next") {
            element.innerHTML = "navigate_before"
            element.classList.add("active")
        } else {
            element.innerHTML = "navigate_next"
            element.classList.remove("active")
        }
    }
}
function arrowClose(element) {
    element.innerHTML = "navigate_next"
    element.classList.remove("active")
}