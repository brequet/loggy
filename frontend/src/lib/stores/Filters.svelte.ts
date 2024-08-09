class Filters {
  appNames = $state<string[]>([]);
  levels = $state<string[]>([]);
  startDate = $state<Date | null>(null);
  endDate = $state<Date | null>(null);

  reset() {
    this.appNames = [];
    this.levels = [];
    this.startDate = null;
    this.endDate = null;
  }
}

export const filters = new Filters();
