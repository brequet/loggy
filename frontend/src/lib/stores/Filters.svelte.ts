class Filters {
  appNames = $state<string[]>([]);
  levels = $state<string[]>([]);
  startDate = $state<Date | null>(null);
  endDate = $state<Date | null>(null);

  queryParams = $derived(
    new URLSearchParams({
      appNames: this.appNames.join(","),
      levels: this.levels.join(","),
      startDate: this.startDate?.toISOString() || "",
      endDate: this.endDate?.toISOString() || "",
    })
  );

  reset() {
    this.appNames = [];
    this.levels = [];
    this.startDate = null;
    this.endDate = null;
  }
}

export const filters = new Filters();
