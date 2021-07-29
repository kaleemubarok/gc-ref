
-- Create table and index for account

CREATE TABLE public.account (
	user_id varchar(32) NOT NULL,
	username varchar(64) NOT NULL,
	"password" varchar(255) NOT NULL,
	salt varchar(32) NOT NULL,
	created_at timestamp(0) NULL,
	profile_pic varchar(255) NULL,
	CONSTRAINT account_pk PRIMARY KEY (user_id),
	CONSTRAINT account_un UNIQUE (username)
);
CREATE INDEX account_username_idx ON public.account USING btree (username);
CREATE INDEX account_user_id_idx ON public.account USING btree (user_id);

INSERT INTO public.account
(user_id, username, "password", salt, created_at, profile_pic)
VALUES  ('ac1jyg1x', 'kevink123', 'cb01ac39710110aa60c078e86be242d0f979a4b42cfdc656e7c753a3f101c2f6', 'IZRgBmyArKCtzkjkZIvaBjMkXVbWGvbq', '2021-04-06 20:38:46.000', ''),
        ('2viin1oy', 'akbar_test', '246add16c1643100d023f7510847e6d40c24d6f5bf5ffaa3ce81ee4652a56ebd', 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2021-04-06 20:48:41.000', ''),
        ('kiw3azqm', 'asdkevin', 'fa5480cfe05cddf946331bbd7a63f8b98bc637f96b6a46ae714447d820561043', 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2021-04-06 21:43:36.000', ''),
        ('4xmq3ipa', 'newbm', 'fa5480cfe05cddf946331bbd7a63f8b98bc637f96b6a46ae714447d820561043', 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2021-04-01 21:08:19.000', ''),
        ('jqtv9t5v', 'bmbmbm', 'fa5480cfe05cddf946331bbd7a63f8b98bc637f96b6a46ae714447d820561043', 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2021-04-01 21:08:19.000', ''),
        ('mq2jytqq', 'kevinqwe', 'fa5480cfe05cddf946331bbd7a63f8b98bc637f96b6a46ae714447d820561043', 'XVlBzgbaiCMRAjWwhTHctcuAxhxKQFDa', '2021-04-24 14:26:33.000', '');
