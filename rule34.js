// ==UserScript==
// @name           Rule34 Copy Photo
// @description    Copy photo
// @author         Maldan
// @include        https://rule34.xxx/*
// @version        1.0
// ==/UserScript==

let PHOTO_DATA = "";
let CURRENT_ELEMENT = null;

const ddddd228 = () => {
  const year = new Date().getFullYear();
  const month = (`00` + (new Date().getMonth() + 1)).slice(-2);
  const date = (`00` + new Date().getDate()).slice(-2);

  return `${year}_${month}_${date}`;
};

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
              path: `/home/maldan/image/rule34/${ddddd228()}/`,
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
