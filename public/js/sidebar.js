function toggleNav() {
    var navToggle = document.getElementById("navToggle");
    var newPosts = document.getElementById("newposts");
    if (newPosts.style.display != "block") {
        newPosts.style.display = "block"
        navToggle.innerHTML = "&and;"
    } else {
        newPosts.style.display = "none"
        navToggle.innerHTML = "&or;"
    }
}
