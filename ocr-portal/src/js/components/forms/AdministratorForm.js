// @flow weak

import React, { Component } from 'react'

import { FormContainer } from '../containers/FormContainer';
import { errorsFor, slugify } from '../../common/util';

import type { APIErrors, APIProject, APIUser, APINewAdministrator, APIAdministrator } from '../../api/types';
import { FKApiClient } from '../../api/api';

type Props = {
project: APIProject,
administrators: APIAdministrator[],

cancelText?: string;
saveText?: ?string;
onCancel?: () => void;
onSave: (e: APINewAdministrator) => Promise<?APIErrors>;
}

export class AdministratorForm extends Component {
    props: Props;
    state: {
    users: APIUser[],
    userId: number,
    saveDisabled: boolean,
    errors: ?APIErrors
    }

    constructor(props: Props) {
        super(props);
        this.state = {
            users: [],
            userId: -1,
            saveDisabled: true,
            errors: null
        }
        this.loadUsers();
    }

    async loadUsers() {
        const usersRes = await FKApiClient.get().getUsers();
        if (usersRes.type === 'ok' && usersRes.payload) {
            const administratorsIds = this.props.administrators.map(administrator => administrator.user_id);
            const availableUsers = usersRes.payload.users.filter(user => administratorsIds.indexOf(user.id) < 0);
            this.setState({
                    users: availableUsers
                } || []);
        }
    }

    handleInputChange(event) {
        const target = event.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        const name = target.name;

        this.setState({
            [name]: value
        });
    }

    async save() {
        const {project} = this.props;
        const {userId} = this.state;
        const errors = await this.props.onSave({
            user_id: parseInt(userId)
        });
        if (errors) {
            this.setState({
                errors
            });
        }
    }

    render() {

        return (
            <FormContainer onSave={ this.save.bind(this) } onCancel={ this.props.onCancel } saveText={ this.props.saveText } cancelText={ this.props.cancelText }>
                <div className="form-group">
                    <label htmlFor="userId">User</label>
                    <select name="userId" className='lg' value={ this.state.userId } onChange={ this.handleInputChange.bind(this) }>
                        <option value={ null }>Select a user</option>
                        { this.state.users.map((user, i) => <option key={ i } value={ user.id }>
                                                                { user.username }
                                                            </option>) }
                    </select>
                    { errorsFor(this.state.errors, 'userId') }
                </div>
            </FormContainer>
        )
    }
}