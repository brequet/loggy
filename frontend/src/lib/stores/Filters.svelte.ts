class Filters {
  appNames = $state<string[]>([]);
  levels = $state<string[]>([]);
  startDate = $state<string | null>(null);
  endDate = $state<string | null>(null);

  activeFilterCount = $derived.by(() => {
    let count = 0;
    if (this.appNames.length > 0) count++;
    if (this.levels.length > 0) count++;
    if (this.startDate !== null) count++;
    if (this.endDate !== null) count++;
    return count;
  })

  reset() {
    this.appNames = [];
    this.levels = [];
    this.startDate = null;
    this.endDate = null;
  }
}

export const filters = new Filters();
