import { Box } from '@suid/material';
import { MemeCard } from '../../components/MemeCard/MemeCard';
import { MemeInput } from '../../components/MemeInput';
import styles from './Home.module.css';

export const Home = () => {
	return (
		<main class={styles['main']}>
			<Box class={styles['meme-container']}>
				<MemeInput />
				<MemeCard />
			</Box>
		</main>
	);
};
