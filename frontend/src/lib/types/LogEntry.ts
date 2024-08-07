export interface LogEntry {
    timestamp: string;
    level: string;
    content: string;
    app_name: string;
    filename: string;
    raw: string;
}
