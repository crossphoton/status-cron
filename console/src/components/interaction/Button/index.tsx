"use client";
import { FC } from "react";
import style from "./button.module.css"

interface ButtonProps {
  text: string;
  onClick?: () => void;
}

const Button: FC<ButtonProps> = ({ text, onClick }) => {
  return (
    <div className={style.button}>
      <div className="button" onClick={onClick}>
        {text}
      </div>
    </div>
  );
};

export default Button;
