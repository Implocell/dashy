export const getPoem = async (text: string) => {
  return await fetch("/api/meme", {
    method: "POST",
    body: JSON.stringify({ text: "Some text" }),
  }).then((res) => res.json());
};

