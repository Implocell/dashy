import { Card, CardContent, CardMedia, Typography } from "@suid/material";
import styles from "./MemeCard.module.css";
export const MemeCard = () => {
  return (
    <Card class={styles["memecard"]}>
      <CardMedia />
      <CardContent class={styles["poem"]}>
        <Typography sx={{ fontSize: 14 }} gutterBottom>
          something
        </Typography>
      </CardContent>
    </Card>
  );
};
