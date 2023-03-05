export const createMeme = async (text: string) =>
	(
		await fetch('/api/meme', {
			method: 'POST',
			body: JSON.stringify({ text }),
		})
	).json();
