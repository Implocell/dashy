import { Meme } from "../types/meme";

export const getMemes = async(): Promise<Meme[]> => (await fetch(`/api/memes`)).json();