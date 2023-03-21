import { Card, CardContent, CardMedia, Typography } from "@suid/material";

import styles from "./MemeCard.module.css";
import errorImg from "../../assets/depressed.png";
import {Meme} from "../../types/meme";



interface Props {
 meme: Meme
}

export const MemeCard = ({meme}: Props) => {
  const imageUrl = meme.url ?? errorImg
  return (
    <div class="slide">
   <img src={meme.url} />
   </div>
  );
};
