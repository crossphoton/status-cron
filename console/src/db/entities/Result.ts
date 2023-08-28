import instance from "../instance";

import { DataTypes } from "sequelize";

instance.define("results", {
    serviceId: {
        type: DataTypes.UUIDV4,
        allowNull: false,
    },
    success: {
        type: DataTypes.BOOLEAN,
        allowNull: false,
        defaultValue: false,
    },
    reason: {
        type: DataTypes.STRING,
        allowNull: true,
    },
    cronTime: {
        type: DataTypes.STRING,
        allowNull: false,
        field: "cron_time",
    },
});