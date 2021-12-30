// ==UserScript==
// @name           Discord Copy Photo
// @description    Copy photo
// @author         Maldan
// @include        https://discord.com/*
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
              path: `/home/maldan/image/discord/${ddddd228()}/`,
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
        CURRENT_ELEMENT.style.filter = "brightness(0.5)";
      }
    } catch {
      alert("Error");
    }
  }
});

setInterval(() => {
  [...document.querySelectorAll("img")].forEach((x) => {
    x.onmouseover = () => {
      PHOTO_DATA = x.getAttribute("src").split("?")[0];
      x.style.opacity = "0.5";
      CURRENT_ELEMENT = x;
    };
    x.onmouseout = () => {
      x.style.opacity = "1";
      PHOTO_DATA = "";
      CURRENT_ELEMENT = null;
    };
  });
}, 1000);
