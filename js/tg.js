// ==UserScript==
// @name           Tg Copy Photo
// @description    Copy photo
// @author         Maldan
// @include        https://web.telegram.org/k/
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
    const photoData = PHOTO_DATA;
    PHOTO_DATA = null;
    try {
      let blob = await fetch(photoData, {
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
              path: `/home/maldan/image/tg/${ddddd228()}`,
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
      } catch {}
    };
    x.onmouseout = () => {
      x.style.opacity = "1";
      PHOTO_DATA = "";
      CURRENT_ELEMENT = null;
    };
  });
}, 1000);
