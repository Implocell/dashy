import { createSignal, onCleanup } from 'solid-js';

const exampleMemeInputs = ['Doffen har dævva...', 'Aldri langt til repern...', 'Like virgin mary...', ''];

export const useMemeExampleText = () => {
	const [memeTextIndex, setMemeTextIndex] = createSignal(0);
	const [memeIndex, setMemeIndex] = createSignal<number>(0);
	const [memeText, setMemeText] = createSignal('');

	const updateMemeTextIndex = () => {
		if (memeTextIndex() >= exampleMemeInputs[memeIndex()].length) {
			setMemeIndex((memeIndex) => (memeIndex + 1) % exampleMemeInputs.length);
			setMemeTextIndex(0);
			setMemeText('');
		} else {
			const newMemeTextIndex = memeTextIndex() + 1;
			setMemeText(exampleMemeInputs[memeIndex()].slice(0, newMemeTextIndex));
			setMemeTextIndex(newMemeTextIndex);
		}
	};
	const memeTypeWriterInterval = setInterval(() => updateMemeTextIndex(), 200);
	onCleanup(() => clearInterval(memeTypeWriterInterval));

	return { memeText };
};
