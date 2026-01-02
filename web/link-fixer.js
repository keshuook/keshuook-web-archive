const path = window.location.pathname.split("/").slice(0, 5).join("/");

window.addEventListener("DOMContentLoaded", () => {
    const aTags = document.getElementsByTagName("a");

    Array.from(aTags).forEach(aTag => {
        const href = aTag.getAttribute("href");
        if(!href) return;
        if(href.startsWith("/")) {
            aTag.href = path+href;
        }
    });

    const iframes = document.getElementsByTagName("iframe");

    Array.from(iframes).forEach(iframe => {
        const src = iframe.getAttribute("src");
        if(!src) return;
        if(src.startsWith("/")) {
            iframe.src = path+src;
        }
    });
})