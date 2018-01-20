import AddIcon from '@sourcegraph/icons/lib/Add'
import HelpIcon from '@sourcegraph/icons/lib/Help'
import Loader from '@sourcegraph/icons/lib/Loader'
import * as H from 'history'
import * as React from 'react'
import { Redirect } from 'react-router'
import { Link } from 'react-router-dom'
import { map } from 'rxjs/operators/map'
import { Subject } from 'rxjs/Subject'
import { Subscription } from 'rxjs/Subscription'
import { currentUser } from '../auth'
import { eventLogger } from '../tracking/eventLogger'
import { observeSavedQueries } from './backend'
import { SavedQuery } from './SavedQuery'
import { SavedQueryCreateForm } from './SavedQueryCreateForm'

interface Props {
    location: H.Location
    isLightTheme: boolean
}

interface State {
    savedQueries: GQL.ISavedQuery[]

    /**
     * Whether the saved query creation form is visible.
     */
    creating: boolean

    loading: boolean
    error?: Error
    user: GQL.IUser | null
}

export class SavedQueries extends React.Component<Props, State> {
    public state: State = {
        savedQueries: [],
        creating: false,
        loading: true,
        user: null,
    }

    private componentUpdates = new Subject<Props>()
    private subscriptions = new Subscription()

    constructor(props: Props) {
        super(props)

        this.subscriptions.add(
            observeSavedQueries()
                .pipe(
                    map(savedQueries => ({
                        savedQueries: savedQueries.sort((a, b) => {
                            if (a.description < b.description) {
                                return -1
                            }
                            if (a.description === b.description && a.index < b.index) {
                                return -1
                            }
                            return 1
                        }),
                        loading: false,
                    }))
                )
                .subscribe(newState => this.setState(newState as State), err => console.error(err))
        )
    }

    public componentDidMount(): void {
        this.subscriptions.add(currentUser.subscribe(user => this.setState({ user })))
    }

    public componentWillReceiveProps(newProps: Props): void {
        this.componentUpdates.next(newProps)
    }

    public componentWillUnmount(): void {
        this.subscriptions.unsubscribe()
    }

    public render(): JSX.Element | null {
        if (this.state.loading) {
            return <Loader />
        }

        const isHomepage = this.props.location.pathname === '/search'

        // If not logged in, redirect to sign in
        if (!this.state.user && !isHomepage) {
            const newUrl = new URL(window.location.href)
            // Return to the current page after sign up/in.
            newUrl.searchParams.set('returnTo', window.location.href)
            return <Redirect to={'/sign-up' + newUrl.search} />
        }

        const savedQueries = this.state.savedQueries.filter(savedQuery => {
            if (isHomepage) {
                return savedQuery.showOnHomepage
            }
            return savedQuery
        })

        return (
            <div className="saved-queries">
                {!isHomepage && (
                    <div>
                        <div className="saved-queries__header">
                            <h2>Saved queries</h2>
                            <span className="saved-queries__center">
                                <button
                                    className="btn btn-link"
                                    onClick={this.toggleCreating}
                                    disabled={this.state.creating}
                                >
                                    <AddIcon className="icon-inline" /> Add new query
                                </button>

                                <a
                                    onClick={this.onDidClickQueryHelp}
                                    className="saved-queries__help"
                                    href="https://about.sourcegraph.com/docs/search/#saved-queries"
                                    target="_blank"
                                >
                                    <small>
                                        <HelpIcon className="icon-inline" />
                                        <span>Help</span>
                                    </small>
                                </a>
                            </span>
                        </div>
                        {this.state.creating && (
                            <SavedQueryCreateForm
                                onDidCreate={this.onDidCreateSavedQuery}
                                onDidCancel={this.toggleCreating}
                            />
                        )}
                        {!this.state.creating &&
                            this.state.savedQueries.length === 0 && <p>You don't have any saved queries yet.</p>}
                    </div>
                )}
                <div>
                    {savedQueries.map((savedQuery, i) => (
                        <SavedQuery
                            hideBottomBorder={i === 0 && savedQueries.length > 1}
                            key={i}
                            savedQuery={savedQuery}
                            onDidDuplicate={this.onDidDuplicateSavedQuery}
                            isLightTheme={this.props.isLightTheme}
                        />
                    ))}
                </div>
                {savedQueries.length === 0 &&
                    this.state.user &&
                    isHomepage && (
                        <div className="saved-query">
                            <Link to="/search/queries">
                                <div className={`saved-query__row`}>
                                    <div className="saved-query__add-query">
                                        <AddIcon className="icon-inline" /> Add a new query to start monitoring your
                                        code.
                                    </div>
                                </div>
                            </Link>
                        </div>
                    )}
            </div>
        )
    }

    private toggleCreating = () => {
        eventLogger.log('SavedQueriesToggleCreating', { queries: { creating: !this.state.creating } })
        this.setState({ creating: !this.state.creating })
    }

    private onDidCreateSavedQuery = () => {
        eventLogger.log('SavedQueryCreated')
        this.setState({ creating: false })
    }

    private onDidDuplicateSavedQuery = () => {
        eventLogger.log('SavedQueryDuplicated')
    }

    private onDidClickQueryHelp = () => {
        eventLogger.log('SavedQueriesHelpButtonClicked')
    }
}

export class SavedQueriesPage extends SavedQueries {
    public componentDidMount(): void {
        super.componentDidMount()
        eventLogger.logViewEvent('SavedQueries')
    }
}
