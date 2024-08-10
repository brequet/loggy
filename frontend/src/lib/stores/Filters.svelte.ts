class Filters {
  appNames = $state<string[]>([]);
  levels = $state<string[]>([]);
  startDate = $state<string | null>(null);
  endDate = $state<string | null>(null);

  reset() {
    this.appNames = [];
    this.levels = [];
    this.startDate = null;
    this.endDate = null;
  }
}

export const filters = new Filters();
