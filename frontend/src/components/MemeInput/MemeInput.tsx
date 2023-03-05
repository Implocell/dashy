import { Send } from '@suid/icons-material';
import { Box, CircularProgress, IconButton, TextField, Typography } from '@suid/material';
import { createSignal } from 'solid-js';
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
	const [isLoading, setIsLoading] = createSignal(false);
	const [input, setInput] = createSignal('');

	const onSubmit = () => {
		if (input().length > 0) {
			setIsLoading(true);
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
					disabled={isLoading()}
					value={input()}
					onChange={(e) => setInput(e.target.value)}
					InputProps={{
						onKeyDown: onKeyDown,
						endAdornment: <MemeInputEndAdornment isLoading={isLoading()} onSubmit={onSubmit} />,
					}}
				/>
			</Box>
			<Box height={16}>
				<Typography variant="caption" fontStyle="italic" color="GrayText">
					{isLoading() ? 'Creating meme...' : memeText}
				</Typography>
			</Box>
		</Box>
	);
};
