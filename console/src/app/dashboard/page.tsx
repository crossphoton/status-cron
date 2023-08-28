import Service from "@/components/statics/Service";
import "../globals.css";
import style from "./dashboard.module.css";

export default function Dashboard() {
  return (
    <div className={[style.verticalFlex, style.center].join(" ")}>
      <div className={style.header}>
        <h1>Dashboard</h1>
      </div>
      <div className={[style.cardContainer, style.verticalFlex].join(" ")}>
        <h3>Your services</h3>
        <div>
          <Service />
        </div>
      </div>
    </div>
  );
}
