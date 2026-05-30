#!/usr/bin/env zsh

load_user_scripts() {
  emulate -L zsh
  setopt null_glob

  local scripts_dir="$HOME/.scripts"
  local script_file
  local success_count=0
  local failure_count=0
  local total_count=0

  if [[ ! -d "$scripts_dir" ]]; then
    print -u2 "load_user_scripts: directory not found: $scripts_dir"
    return 0
  fi

  local script_files=("$scripts_dir"/*.sh(.N))

  if (( ${#script_files[@]} == 0 )); then
    print "load_user_scripts: no .sh files found in $scripts_dir"
    print "load_user_scripts: completed, success=0, failed=0"
    return 0
  fi

  for script_file in "${script_files[@]}"; do
    (( total_count++ ))
    print "load_user_scripts: sourcing [$total_count/${#script_files[@]}] $script_file"

    if source "$script_file"; then
      (( success_count++ ))
      print "load_user_scripts: loaded successfully: $script_file"
    else
      local exit_code=$?
      (( failure_count++ ))
      print -u2 "load_user_scripts: failed to source: $script_file (exit code: $exit_code)"
    fi
  done

  print "load_user_scripts: completed, success=$success_count, failed=$failure_count"

  if (( failure_count > 0 )); then
    return 1
  fi
  return 0
}

print_load_user_scripts_install_help() {
  emulate -L zsh

  local script_path="${(%):-%x}"
  local target_file="${1:-$HOME/.zshrc}"

  print "Add the following line to $target_file:"
  print "source ${(q)script_path}"
  print ""
  print "Apply it in the current shell:"
  print "source ${(q)target_file}"
  print ""
  print "Verify it in a new zsh session:"
  print "zsh -lic 'type load_user_scripts && load_user_scripts'"
  print ""
  print "If you prefer .zsh_profile, run:"
  print "print 'source ${(q)script_path}' >> ~/.zsh_profile"
}
