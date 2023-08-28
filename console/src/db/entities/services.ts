import instance from "../instance"
import { DataTypes } from "sequelize";

enum ServiceType {
    HTTP= "http",
    REDIS = "redis",
    SQL = "sql",
    MONGO = "mongo",
}

instance.define("services", {
    id: {
        primaryKey: true,
        type: DataTypes.UUIDV4,
    },
    name: {
        type: DataTypes.STRING,
        allowNull: false,
    },
    url: {
        type: DataTypes.STRING,
        allowNull: false,
    },
    type: {
        type: DataTypes.ENUM(ServiceType.HTTP, ServiceType.REDIS, ServiceType.SQL, ServiceType.MONGO),
        allowNull: false,
    },
    cron: {
        type: DataTypes.STRING,
        allowNull: false
    },
    data: {
        type: DataTypes.JSON,
        allowNull: false,
    },
});