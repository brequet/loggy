import type { LogEntriesResult } from "../types/LogEntry";

const API_URL = "http://localhost:8080/api";

export async function fetchLogEntriesResult(queryParams: string): Promise<LogEntriesResult> {
  const response = await fetch(`${API_URL}/logs?${queryParams}`);
  if (!response.ok) {
    throw new Error("Failed to fetch logs");
  }
  return response.json();
}

export async function fetchAppNames(): Promise<string[]> {
  const response = await fetch(`${API_URL}/apps`);
  if (!response.ok) {
    throw new Error("Failed to fetch apps");
  }
  return response.json();
}
