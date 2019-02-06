function toggleNav() {
    var sidebarContent = document.getElementById("sidebarContent");
    var navToggle = document.getElementById("navToggle");
    if (sidebarContent.style.display != "block") {
        sidebarContent.style.display = "block"
        navToggle.innerHTML = "&and;"
    } else {
        sidebarContent.style.display = ""
        navToggle.innerHTML = "&or;"
    }
}
