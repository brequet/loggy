export interface LogEntry {
    timestamp: string;
    level: string;
    content: string;
    app_name: string;
    filename: string;
    raw: string;
}

export interface LogEntriesResult {
    entries: LogEntry[];
    total_count: number;
    page: number;
    page_size: number;
}

export enum LogLevel {
    DEBUG = "DEBUG",
    INFO = "INFO",
    WARN = "WARN",
    ERROR = "ERROR",
    FATAL = "FATAL",
}

export const allLogLevels: LogLevel[] = Object.values(LogLevel);
