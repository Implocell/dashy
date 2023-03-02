import { Box, TextField, Typography } from '@suid/material';
import { useMemeExampleText } from './hooks';

export const MemeInput = () => {
	const { memeText } = useMemeExampleText();

	return (
		<Box pb={4} sx={{ display: 'flex' }} flexDirection="column">
			<TextField size="small" placeholder="insert meme here!" />
			<Box height={16}>
				<Typography variant="caption" fontStyle="italic" color="GrayText">
					{memeText}
				</Typography>
			</Box>
		</Box>
	);
};
