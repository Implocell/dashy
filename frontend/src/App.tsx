import type { Component } from 'solid-js';

import { CssBaseline, ThemeProvider } from '@suid/material';
import { theme } from './theme';
import { Route, Router, Routes } from '@solidjs/router';
import { Home } from './pages/Home';

const App: Component = () => {
	return (
		<Router>
			<ThemeProvider theme={theme}>
				<CssBaseline />
				<Routes>
					<Route path="/" component={Home} />
				</Routes>
			</ThemeProvider>
		</Router>
	);
};

export default App;
