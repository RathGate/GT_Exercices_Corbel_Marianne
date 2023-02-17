// *RESPONSIVE NAVBAR:

document.querySelector(".icon-menu").addEventListener("click", function() { 
    document.querySelector(".nav-link-container").classList.toggle("visible")
})

document.querySelectorAll(".nav-link").forEach(element => {
    element.addEventListener("click", function() { 
        if (window.innerWidth < 750) {
            document.querySelector(".nav-link-container").classList.remove("visible")
        }
    })
});

// On switch from phone view to big screen view, removes the "visible" class on 
// the menu (needed only when using burger menu)
window.addEventListener("resize", function() {
    if (this.innerWidth > 950) {
        this.document.querySelector(".nav-link-container").classList.remove("visible")
    } 
})

console.log("bonjour")