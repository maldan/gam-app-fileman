// ==UserScript==
// @name           Rule34 Copy Photo
// @description    Copy photo
// @author         Maldan
// @include        https://rule34.xxx/*
// @version        1.0
// ==/UserScript==

let PHOTO_DATA = "";
let CURRENT_ELEMENT = null;

document.addEventListener("keydown", async (e) => {
  if (e.key === "1" && PHOTO_DATA) {
    try {
      const r = await fetch(
        `http://192.168.1.94:4000/api/main/applicationData`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            appId: "maldan/fileman",
            data: JSON.stringify({
              type: "download_photo",
              path: `/home/maldan/image/rule34/${new Date().getFullYear()}_${new Date().getMonth()}_${new Date().getDate()}`,
              url: PHOTO_DATA,
            }),
          }),
        }
      );
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
  [...document.querySelectorAll("img")].forEach((x) => {
    x.onmouseover = () => {
      try {
        PHOTO_DATA = x.getAttribute("src");
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
