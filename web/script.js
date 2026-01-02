function getDate(factor, lastDate, firstDate) {
    return new Date(firstDate.valueOf() + Math.round((factor)*(lastDate - firstDate)));
}

(async function () {
    const firstReq = await fetch("/api/firstcommit/");
    const firstDate = new Date(await firstReq.text());
    const lastReq = await fetch("/api/lastcommit/");
    const lastDate = new Date(await lastReq.text());

    const iframe = document.getElementById("display");
    const waybackContainer = document.getElementById("wayback-nav");
    waybackContainer.style.height = "140px";
    waybackContainer.style.padding = "10px";

    const timeSlider = document.getElementById("timeSlider");
    const labels = document.getElementsByClassName("year-label");
    const dateSelector = document.getElementById("date-selector");
    const urlElement = document.getElementById("current-url");
    
    const firstYear = firstDate.getFullYear();
    const lastYear = lastDate.getFullYear();
    labels[0].innerHTML = firstYear
    labels[1].innerHTML =  Math.round(firstYear + (lastYear - firstYear)/3)
    labels[2].innerHTML =  Math.round(firstYear + 2*(lastYear - firstYear)/3)
    labels[3].innerHTML = lastYear

    dateSelector.value = (new Date()).toISOString().split('T')[0];
    iframe.src = `/wayback/${dateSelector.value.replace(/-/g, "/")}/`;

    dateSelector.addEventListener("input", () => {
        iframe.src = `/wayback/${dateSelector.value.replace(/-/g, "/")}/`;
    })

    timeSlider.addEventListener("input", () => {
        const selectedDate = getDate(timeSlider.value/timeSlider.max, lastDate, firstDate);

        dateSelector.value = selectedDate.toISOString().split('T')[0];

        urlElement.href = `/wayback/${selectedDate.getFullYear()}/${selectedDate.getMonth() + 1}/${selectedDate.getDate()}/`;
    });
    
    timeSlider.addEventListener("change", () => {
        const selectedDate = getDate(timeSlider.value/timeSlider.max, lastDate, firstDate);

        iframe.src = `/wayback/${selectedDate.getFullYear()}/${selectedDate.getMonth() + 1}/${selectedDate.getDate()}/`;
    });

    let prevY = 0;

    iframe.addEventListener("load", () => {
        try {
            const currentURL = iframe.contentWindow.location.href;
        } catch (e) {
            alert("You've left the webarchive!");
            urlElement.href = "";
            urlElement.innerHTML = "Not in Web Archive"
        }

        urlElement.href = iframe.contentWindow.location.href;
        urlElement.innerHTML = iframe.contentWindow.location.href.split("/").slice(3).join("/");
        waybackContainer.style.height = "140px";
        waybackContainer.style.padding = "10px";
        iframe.contentWindow.addEventListener("scroll", () => {
            const y = iframe.contentWindow.scrollY
            if(prevY - y < -2) {
                waybackContainer.style.height = "0px";
                waybackContainer.style.padding = "";
            }
            if (prevY - y > 50 ||  y == 0) {
                waybackContainer.style.height = "140px";
                waybackContainer.style.padding = "10px";
            } 
            prevY = y;  
        });
    });
})()