[...document.querySelectorAll("a[aria-label=фотография]")]
  .map((x) =>
    JSON.parse(
      x
        .getAttribute("onclick")
        .replace(/return showPhoto\('.*', '.*', (.*?), event\);?/gm, "$1")
    )
  )
  .map((x) => x.temp.w || x.temp.z || x.temp.y || x.temp.x)
  .join(" ");

// ==UserScript==
// @name           Vk Copy Photo
// @description    Copy photo
// @author         Maldan
// @include        https://vk.com/feed*
// @version        1.0
// ==/UserScript==

let PHOTO_DATA = "";
let CURRENT_ELEMENT = null;

document.addEventListener("keydown", async (e) => {
  if (e.key === "1" && PHOTO_DATA) {
    try {
      const r = await fetch(`http://127.0.0.1:16001/api/main/applicationData`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          appId: "maldan/fileman",
          data: JSON.stringify({
            type: "download_photo",
            path: `/home/maldan/image/vk/${new Date().getFullYear()}_${new Date().getMonth()}_${new Date().getDate()}`,
            url: PHOTO_DATA,
          }),
        }),
      });
      if (!r.ok) {
        alert("Error");
        return;
      }
      await r.json();
      if (CURRENT_ELEMENT) {
        CURRENT_ELEMENT.style.border = "1px solid #00ff00";
      }
    } catch {
      alert("Error");
    }
  }
});

setInterval(() => {
  [...document.querySelectorAll("a[aria-label]")].forEach((x) => {
    x.onmouseover = () => {
      try {
        const y = JSON.parse(
          x
            .getAttribute("onclick")
            .replace(/return showPhoto\('.*', '.*', (.*?), event\);?/gm, "$1")
        );
        PHOTO_DATA = y.temp.w || y.temp.z || y.temp.y || y.temp.x;
        x.style.border = "1px solid #fe0000";
        CURRENT_ELEMENT = x;
      } catch {}
    };
    x.onmouseout = () => {
      x.style.border = "none";
      PHOTO_DATA = "";
      CURRENT_ELEMENT = null;
    };
  });
}, 1000);
