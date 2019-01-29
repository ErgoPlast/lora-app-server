-- +migrate Up
alter table device_keys
    add column gen_app_key bytea not null default decode('00000000000000000000000000000000', 'hex');

alter table device_keys
    alter column gen_app_key drop default;

alter table multicast_group
    add column mc_key bytea not null default decode('00000000000000000000000000000000', 'hex'),
    add column f_cnt bigint not null default 0;

alter table multicast_group
    alter column mc_key drop default,
    alter column f_cnt drop default;

create table remote_multicast_setup (
    dev_eui bytea not null references device on delete cascade,
    multicast_group_id uuid not null references multicast_group on delete cascade,
    created_at timestamp with time zone not null,
    updated_at timestamp with time zone not null,
    mc_group_id smallint not null,
    mc_addr bytea not null,
    mc_key_encrypted bytea not null,
    min_mc_f_cnt bigint not null,
    max_mc_f_cnt bigint not null,
    state varchar(20) not null,
    state_provisioned bool not null default false,
    retry_after timestamp with time zone not null,
    retry_count smallint not null,

    primary key(dev_eui, multicast_group_id)
);

create index idx_remote_multicast_setup_state_provisioned on remote_multicast_setup(state_provisioned);
create index idx_remote_multicast_setup_retry_after on remote_multicast_setup(retry_after);

-- +migrate Down
alter table multicast_group
    drop column mc_key,
    drop column f_cnt;

alter table device_keys
    drop column gen_app_key;

drop index idx_remote_multicast_setup_retry_after;
drop index idx_remote_multicast_setup_state_provisioned;
drop table remote_multicast_setup;
