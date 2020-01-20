import React from 'react'
import { SiteAdminRepositoriesOverviewCard } from '../repositories/SiteAdminRepositoriesOverviewCard'
import { SiteAdminUsersOverviewCard } from '../users/SiteAdminUsersOverviewCard'
import { SiteAdminUsageStatisticsOverviewCard } from '../usageStatistics/SiteAdminUsageStatisticsOverviewCard'

export interface SiteAdminOverviewComponent {
    component: React.ComponentType
    noCardClass?: boolean
    fullWidth?: boolean
}

/**
 * Additional components to render on the SiteAdminOverviewPage.
 */
export const siteAdminOverviewComponents: readonly SiteAdminOverviewComponent[] = [
    { component: SiteAdminRepositoriesOverviewCard },
    { component: SiteAdminUsersOverviewCard },
    { component: SiteAdminUsageStatisticsOverviewCard },
]
