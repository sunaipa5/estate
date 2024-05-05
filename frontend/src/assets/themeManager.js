export function themeIcon(){
    let darkIcon = document.getElementById("dark");
    let lightIcon = document.getElementById("light");
    let themeType = localStorage.getItem("theme");

    if (themeType == "light") {
      lightIcon.style.display = "none";
      darkIcon.style.display = "block";
    } else if (themeType == "dark") {
      darkIcon.style.display = "none";
      lightIcon.style.display = "block";
    }
}

