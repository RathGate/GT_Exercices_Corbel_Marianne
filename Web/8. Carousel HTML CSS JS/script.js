var currSlideIndex = 0;
var carouselBody = document.querySelector(".carousel-body")
var slideContainer = document.querySelector(".slide-container")
var slideItem = document.querySelectorAll(".slide-item")
var slideIcons = Array.from(document.querySelector(".slide-icons").children)

var switchAllowed = true

document.querySelector(".prev").addEventListener("click", function() {switchSlide(-1)})
document.querySelector(".next").addEventListener("click", function() {switchSlide(1)})
slideContainer.addEventListener("transitionend", function() {
    console.log("wtf")
    switchAllowed = true
})
function switchSlide(direction) {
    if (!switchAllowed) {
        return
    }
    if (direction == 1 && currSlideIndex != slideItem.length - 1) {
        slideContainer.style.left = slideContainer.offsetLeft - 1000 + "px" 
        currSlideIndex++
        clearSlideIcons(currSlideIndex)
        switchAllowed = false;
    }
    if (direction == -1 && currSlideIndex != 0) {
        slideContainer.style.left = slideContainer.offsetLeft + 1000 + "px"
        currSlideIndex--
        clearSlideIcons(currSlideIndex)
        switchAllowed = false;
    }
    
}

function clearSlideIcons(currentIndex) {
    if (slideIcons.length != 0) {
        slideIcons.forEach(element => {
            element.classList.remove("icon-active")
        });
        slideIcons[currentIndex].classList.add("icon-active")
    }
}

function shortcutSwitch(iconID) {
    if (iconID > slideItem.length -1) {
        return
    }

    if (iconID == currSlideIndex || !switchAllowed) {
        return
    }

    console.log(iconID * 1000)
    slideContainer.style.left = "-" +(iconID * 1000) + "px"
    currSlideIndex = iconID
    clearSlideIcons(currSlideIndex)
    switchAllowed = false;
}


for (let i = 0; i < slideIcons.length; i++) {
    slideIcons[i].addEventListener("click", function() {
        console.log("wtf")
        shortcutSwitch(i)
    })
}
