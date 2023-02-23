export const getPoem = async (text: string) => {
  fetch("/api/meme", {
    method: "POST",
    body: JSON.stringify({ text: "Some text" }),
  })
    .then((res) => {
      return res.json();
    })
    .then((res) => console.log(res));
};
