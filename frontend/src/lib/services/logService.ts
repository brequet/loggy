import type { LogEntry } from "../types/LogEntry";

const API_URL = "http://localhost:8080/api";

export async function fetchLogs(queryParams: string): Promise<LogEntry[]> {
  const response = await fetch(`${API_URL}/logs?${queryParams}`);
  if (!response.ok) {
    throw new Error("Failed to fetch logs");
  }
  return response.json();
}
