import { Send } from '@suid/icons-material';
import { Box, CircularProgress, IconButton, TextField, Typography } from '@suid/material';
import { createResource, createSignal } from 'solid-js';
import { createMeme } from '../../api/createMeme';
import { useMemeExampleText } from './hooks';

interface MemeInputEndAdornmentProps {
	isLoading: boolean;
	onSubmit: () => void;
}

const MemeInputEndAdornment = (props: MemeInputEndAdornmentProps) => {
	return (
		<IconButton size="small" onClick={() => props.onSubmit()} disabled={props.isLoading}>
			{props.isLoading ? <CircularProgress size={24} /> : <Send color="primary" fontSize="small" />}
		</IconButton>
	);
};

export const MemeInput = () => {
	const { memeText } = useMemeExampleText();
	const [input, setInput] = createSignal<string | undefined>(undefined);
	const [resourceInput, setResourceInput] = createSignal<string | undefined>(undefined);
	const [data] = createResource(resourceInput, createMeme);

	const onSubmit = () => {
		const inputValue = input();
		if (typeof inputValue === 'string' && inputValue.length > 0) {
			setResourceInput(inputValue);
		}
	};

	const onKeyDown = (event: KeyboardEvent) => {
		if (event.key === 'Enter') {
			onSubmit();
		}
	};

	return (
		<Box pb={4} sx={{ display: 'flex' }} flexDirection="column">
			<Box>
				<TextField
					size="small"
					placeholder="Insert meme here!"
					disabled={data.loading}
					value={input()}
					onChange={(e) => setInput(e.target.value)}
					InputProps={{
						onKeyDown: onKeyDown,
						endAdornment: <MemeInputEndAdornment isLoading={data.loading} onSubmit={onSubmit} />,
					}}
				/>
			</Box>
			<Box height={16}>
				<Typography variant="caption" fontStyle="italic" color="GrayText">
					{data.loading ? 'Creating meme...' : memeText}
				</Typography>
			</Box>
			<Box>{JSON.stringify(data())}</Box>
		</Box>
	);
};
