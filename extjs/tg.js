// ==UserScript==
// @name           Tg Copy Photo
// @description    Copy photo
// @author         Maldan
// @include        https://web.telegram.org
// @version        1.0
// ==/UserScript==

let PHOTO_DATA = "";
let CURRENT_ELEMENT = null;

function blobToBase64(blob) {
  return new Promise((resolve) => {
    var reader = new FileReader();
    reader.readAsDataURL(blob);
    reader.onloadend = function () {
      let x = reader.result;
      resolve(x.substr(x.indexOf(",") + 1));
    };
  });
}

document.addEventListener("keydown", async (e) => {
  if (e.key === "1" && PHOTO_DATA) {
    try {
      let blob = await fetch(PHOTO_DATA, {
        method: "GET",
      });
      blob = await blob.blob();

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
              type: "download_blob_photo",
              path: `/home/maldan/image/tg/${new Date().getFullYear()}_${new Date().getMonth()}_${new Date().getDate()}`,
              data: await blobToBase64(blob),
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
  [...document.querySelectorAll("img.media-photo")].forEach((x) => {
    x.onmouseover = () => {
      try {
        PHOTO_DATA = x.getAttribute("src");
        x.style.opacity = "0.5";
        CURRENT_ELEMENT = x;
        console.log(PHOTO_DATA);
      } catch {}
    };
    x.onmouseout = () => {
      x.style.opacity = "1";
      PHOTO_DATA = "";
      CURRENT_ELEMENT = null;
    };
  });
}, 1000);
